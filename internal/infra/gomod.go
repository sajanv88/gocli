package infra

func InitGoModule(dir, modulePath string) error {
	if err := EnsureDir(dir); err != nil {
		return err
	}
	return Run(dir, "go", "mod", "init", modulePath)
}
