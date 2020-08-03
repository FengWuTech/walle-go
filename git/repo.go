package git

import (
	"fmt"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
	"walle-go/pkg/fileutil"
)

const ModePerm = 0755

type Repo struct {
	Path     string
	AuthInfo *ssh.PublicKeys
}


// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	//os.Exit(1)
}


//----------------------------------


//判断是否为git目录
func  (r *Repo)isGitDir() bool{
	d := r.Path + "/.git"

	if fileutil.IsDir(d) {
		if fileutil.IsDir(path.Join(d, "objects")) && fileutil.IsDir(path.Join(d, "refs")) {
			headRef := path.Join(d, "HEAD")

			//1是文件
			if fileutil.IsFile(headRef){
				return true
			}
			//2是指定条件的链接
			realName, err := os.Readlink(headRef)
			if err!=nil{
				return false
			}
			return strings.HasPrefix(realName, "refs")

		} else if fileutil.IsFile(path.Join(d, "gitdir")) && fileutil.IsFile(path.Join(d, "commondir")) && fileutil.IsFile(path.Join(d, "gitfile")) {
			return false
		}
	}
	return false
}

// 创建
func  (r *Repo) new() {
	currentUser, _ := user.Current()
	sshAuth, _ := ssh.NewPublicKeysFromFile("git", currentUser.HomeDir+"/.ssh/id_rsa", "")
	r.AuthInfo = sshAuth
}

func  (r *Repo) Init(url string, branch string) {
	if branch == "" {
		branch = "master"
	}
	//创建这个目录
	if !fileutil.Exists(r.Path){
		err:=os.MkdirAll(r.Path, ModePerm)
		if err!=nil{
			return
		}
		//切换下权限
		os.Chmod(r.Path, ModePerm)
	}
	if r.isGitDir() {
		// git pull
		r.Checkout2Branch(branch)
		r.Pull()
	} else {
		// git clone
		r.Clone(url)
	}
}

//检出项目
func  (r *Repo)Clone(url string) {
	r.new()
	_, err := gogit.PlainClone(r.Path, false, &gogit.CloneOptions{
		URL:      url,
		//RemoteName:"origin",
		//ReferenceName:"master",
		Progress: os.Stdout,
		Auth: r.AuthInfo,
	})
	CheckIfError(err)
}

//更新项目
func  (r *Repo)Pull() {
	r.new()
	// We instantiate a new repository targeting the given path (the .git folder)
	repo, err := gogit.PlainOpen(r.Path)
	CheckIfError(err)

	// Get the working directory for the repository
	w, err := repo.Worktree()
	CheckIfError(err)

	err = w.Pull(&gogit.PullOptions{
		//RemoteName: "origin",
		Auth: r.AuthInfo,
	})
	CheckIfError(err)
}

//切换到某个分支
func  (r *Repo)Checkout2Branch(branch string) {
	r.new()
	// We instantiate a new repository targeting the given path (the .git folder)
	repo, err := gogit.PlainOpen(r.Path)
	CheckIfError(err)

	// Get the working directory for the repository
	w, err := repo.Worktree()
	CheckIfError(err)

	// ... checking out to branch
	//得保证这个分支存在.
	/*
	ref, err := repo.Head()
	CheckIfError(err)
	ref.Hash()
	 */
	// First try to checkout branch
	err = w.Checkout(&gogit.CheckoutOptions{Create: false, Force: false, Branch: plumbing.NewBranchReferenceName(branch)} )

	if err != nil {
		cmd := exec.Command("git", "checkout", "--track", "origin/" + branch)
		cmd.Dir = r.Path
		cmd.Start()
		err = cmd.Wait()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}
	}
	CheckIfError(err)
}

//切换分支的某个commit
func  (r *Repo)Checkout2Commit(branch string, commit string) {
	r.new()
	r.Checkout2Branch(branch)

	// We instantiate a new repository targeting the given path (the .git folder)
	repo, err := gogit.PlainOpen(r.Path)
	CheckIfError(err)

	// Get the working directory for the repository
	_, err = repo.Worktree()
	CheckIfError(err)

	// ... checking out to branch
	// hash, err := repo.ResolveRevision(plumbing.Revision(commit))
	// 解析不了短sha commit号，无语了
	//CheckIfError(err)
	Info("git checkout %s", commit)
	//err = w.Checkout(&gogit.CheckoutOptions{
	//	Hash: *hash,
	//})
	//CheckIfError(err)
	cmd := exec.Command("git", "checkout", commit)
	cmd.Dir = r.Path
	cmd.Start()
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

//todo
//切换到tag
func  (r *Repo)Checkout2Tag(tag string) {
	// We instantiate a new repository targeting the given path (the .git folder)
	repo, err := gogit.PlainOpen(r.Path)
	CheckIfError(err)

	// Get the working directory for the repository
	w, err := repo.Worktree()
	CheckIfError(err)


	tInfo,err:=repo.Tag(tag)
	CheckIfError(err)

	// ... checking out to branch
	Info("git checkout %s", tag)
	err = w.Checkout(&gogit.CheckoutOptions{
		//Hash: tInfo.Hash(),
		Branch: tInfo.Name(),
	})
	CheckIfError(err)
}

//获取rmote的所有分支
func remoteBranches(s storer.ReferenceStorer) (storer.ReferenceIter, error) {
	refs, err := s.IterReferences()
	if err != nil {
		return nil, err
	}

	return storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsRemote()
	}, refs), nil
}

//获取所有分支
func  (r *Repo)Branches() []interface{}{
	/*
	去除origin/HEAD- > 当前指向
	去除远端前缀
	*/
	r.new()
	// We instantiate a new repository targeting the given path (the .git folder)
	repo, err := gogit.PlainOpen(r.Path)
	CheckIfError(err)

	//---------

	branchList := make([]interface{},0)

	//https://github.com/src-d/go-git/issues/601
	/*
		branchRefs, err := repo.Branches()
	*/
	branchRefs, err := remoteBranches(repo.Storer)
	CheckIfError(err)
	prefix := "refs/remotes/origin/"

	err = branchRefs.ForEach(func(ref *plumbing.Reference) error {
		fmt.Println(ref.Name())

		rawName := string(ref.Name())
		if strings.HasPrefix(rawName, prefix) && !strings.Contains(rawName, "origin/HEAD") {
			branchName := rawName[len(prefix):]
			branchList = append(branchList, branchName)
		}
		return nil
	})
	CheckIfError(err)

	return branchList
}


//获取所有tag，按时间倒序
func  (r *Repo)Tags()[]string {
	// We instantiate a new repository targeting the given path (the .git folder)
	repo, err := gogit.PlainOpen(r.Path)
	CheckIfError(err)

	// List all tag references, both lightweight tags and annotated tags
	Info("git show-ref --tag")

	tagList := make([]string,0)

	tagrefs, err := repo.Tags()
	CheckIfError(err)
	err = tagrefs.ForEach(func(ref *plumbing.Reference) error {
		obj, err := repo.TagObject(ref.Hash())
		if err!=nil{
			return err
		}
		tagList = append(tagList,obj.Name)
		return nil
	})
	CheckIfError(err)

	return tagList
}



//获取分支的commits
func  (r *Repo)Commits(branch string) []map[string]interface{} {
	r.Checkout2Branch(branch)

	// We instantiate a new repository targeting the given path (the .git folder)
	repo, err := gogit.PlainOpen(r.Path)
	CheckIfError(err)

	commitList := make([]map[string]interface{},0)

	tagrefs, err := repo.Log(&gogit.LogOptions{All:true})
	CheckIfError(err)
	err = tagrefs.ForEach(func(t *object.Commit) error {

		item := map[string]interface{}{
			"id":t.ID().String()[:10],
			"message":t.Message,
			"name":t.Author.Name,
		}
		commitList = append(commitList,item)

		return nil
	})
	CheckIfError(err)

	return commitList[:20]
}