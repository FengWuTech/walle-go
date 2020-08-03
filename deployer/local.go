package deployer

import (
	"fmt"
	"strings"
	"walle-go/git"
	"walle-go/logger"
)

func (d *Deployer) preDeploy() error {
	d.Stage = StagePrevDeploy
	d.Sequence = 1

	d.InitRepo()
	d.Localhost.Dir = d.DirCodebaseProject
	if *d.Project.PrevDeploy != "" {
		commands := strings.Split(*d.Project.PrevDeploy, "\n")
		for _, commandStr := range commands {
			err := d.Localhost.Run(d.config(), commandStr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Deployer) deploy() error {
	d.Stage = StageDeploy
	d.Sequence = 2

	d.Localhost.Dir = CodeBase
	commandStr := fmt.Sprintf("cp -rf %s %s", d.DirCodebaseProject, d.ReleaseVersion)
	d.Localhost.Run(d.config(), commandStr)
	logger.Debugf("cd %s command: %s", CodeBase, fmt.Sprintf("cp -rf %s %s", d.DirCodebaseProject, d.ReleaseVersion))
	path := fmt.Sprintf("%s%s", CodeBase, d.ReleaseVersion)
	repo := git.Repo{Path:path}
	logger.Debugf("cd %s command git checkout to branch %s commit %s", path, *d.Task.Branch, *d.Task.CommitId)
	repo.Checkout2Commit(*d.Task.Branch, *d.Task.CommitId)
	return nil
}

func (d *Deployer) postDeploy() error {
	d.Stage = StagePostDeploy
	d.Sequence = 3

	d.Localhost.Dir = d.DirReleaseVersion
	if *d.Project.PostDeploy != "" {
		commands := strings.Split(*d.Project.PostDeploy, "\n")
		for _, commandStr := range commands {
			d.Localhost.Run(d.config(), commandStr)
		}
	}

	// 打包
	files := ""
	if *d.Project.IsInclude == 1 {
		files = IncludesFormat(d.ReleaseVersion, *d.Project.Excludes)
	} else {
		files = ExcludesFormat(d.ReleaseVersion, *d.Project.Excludes)
	}
	d.Localhost.Dir = CodeBase
	commandStr := fmt.Sprintf("tar zcf %s%s.tgz %s", CodeBase, d.ReleaseVersion, files)
	d.Localhost.Run(d.config(), commandStr)
	return nil
}
