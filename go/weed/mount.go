package main

type MountOptions struct {
	filer *string
	dir   *string
}

var (
	mountOptions MountOptions
)

func init() {
	cmdMount.Run = runMount // break init cycle
	cmdMount.IsDebug = cmdMount.Flag.Bool("debug", false, "verbose debug information")
	mountOptions.filer = cmdMount.Flag.String("filer", "localhost:8888", "dfs filer location")
	mountOptions.dir = cmdMount.Flag.String("dir", ".", "mount dfs filer to this directory")
}

var cmdMount = &Command{
	UsageLine: "mount -filer=localhost:8888 -dir=/some/dir",
	Short:     "mount dfs filer to a directory as file system in userspace(FUSE)",
	Long: `mount dfs filer to userspace.

  Pre-requisites:
  1) have DFS master and volume servers running
  2) have a "dfs filer" running
  These 2 requirements can be achieved with one command "dfs server -filer=true"

  This uses bazil.org/fuse, whichenables writing FUSE file systems on
  Linux, and OS X.

  On OS X, it requires OSXFUSE (http://osxfuse.github.com/).

  `,
}
