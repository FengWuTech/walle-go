package deployer

import (
	"fmt"
	"strings"
	"walle-go/waller"
)

func (d *Deployer) preRelease(w *waller.Waller) error {
	d.Stage = StagePrevRelease
	d.Sequence = 4

	dirTargetReleases := *d.Project.TargetReleases
	_, err := w.RemoteRun(d.config(), fmt.Sprintf("mkdir -p %s", dirTargetReleases))
	if err != nil {
		return err
	}

	err = w.Put(d.config(), fmt.Sprintf("%s%s.tgz", CodeBase, d.ReleaseVersion), fmt.Sprintf("%s/%s.tgz", dirTargetReleases, d.ReleaseVersion))
	if err != nil {
		return err
	}

	_, err = w.RemoteRun(d.config(), fmt.Sprintf("cd %s && tar zxf %s.tgz", dirTargetReleases, d.ReleaseVersion))
	if err != nil {
		return err
	}

	err = d.preReleaseCustom(w)
	if err != nil {
		return err
	}
	return nil
}

func (d *Deployer) preReleaseCustom(w *waller.Waller) error {
	dirTargetReleases := *d.Project.TargetReleases
	if *d.Project.PrevRelease != "" {
		commands := strings.Split(*d.Project.PrevRelease, "\n")
		for _, commandStr := range commands {
			w.RemoteRun(d.config(), fmt.Sprintf("cd %s/%s;%s", dirTargetReleases, d.ReleaseVersion, commandStr))
		}
	}
	return nil
}

func (d *Deployer) release(w *waller.Waller) error {
	d.Stage = StageRelease
	d.Sequence = 5

	dirTargetReleases := *d.Project.TargetReleases
	dirTargetRoot := *d.Project.TargetRoot
	dirCurrentTemp := fmt.Sprintf("current-tmp-%s", d.TaskId)

	// find last release
	command := fmt.Sprintf(`[ -L %s ] && readlink %s || echo ""`, dirTargetRoot, dirTargetRoot)
	output, err := w.RemoteRun(d.config(), command)
	if err != nil {
		return err
	}
	l := strings.Split(output, "/")
	d.LastReleaseVersion = strings.TrimSpace(l[len(l) - 1])

	// create temp dir link
	command = fmt.Sprintf("ln -sfn %s/%s %s/%s", dirTargetReleases, d.ReleaseVersion, dirTargetReleases, dirCurrentTemp)
	_, err = w.RemoteRun(d.config(), command)
	if err != nil {
		return err
	}

	// mv temp link to webroot
	command = fmt.Sprintf("mv -fT %s/%s %s", dirTargetReleases, dirCurrentTemp, dirTargetRoot)
	_, err = w.RemoteRun(d.config(), command)
	if err != nil {
		return err
	}

	return nil
}

func (d *Deployer) postRelease(w *waller.Waller) error {
	d.Stage = StagePostRelease
	d.Sequence = 6

	dirTargetRoot := *d.Project.TargetRoot
	if *d.Project.PostRelease != "" {
		commands := strings.Split(*d.Project.PostRelease, "\n")
		for _, commandStr := range commands {
			_, err := w.RemoteRun(d.config(), fmt.Sprintf("cd %s; %s", dirTargetRoot, commandStr))
			if err != nil {
				return err
			}
		}
	}

	err := d.cleanupRemote(w)
	if err != nil {
		return err
	}
	return nil
}

func (d *Deployer) cleanupRemote(w *waller.Waller) error {
	dirTargetReleases := *d.Project.TargetReleases

	command := fmt.Sprintf("rm -rf %s_%s_*.tgz", d.ProjectId, d.TaskId)
	_, err := w.RemoteRun(d.config(), fmt.Sprintf("cd %s && %s", dirTargetReleases, command))
	if err != nil {
		return err
	}

	command = fmt.Sprintf(`find ./ -name "%s_*" -print | ls -t | tail -n +%d | xargs rm -rf`, d.ProjectId, *d.Project.KeepVersionNum)
	_, err = w.RemoteRun(d.config(), fmt.Sprintf("cd %s && %s", dirTargetReleases, command))
	if err != nil {
		return err
	}

	return nil
}