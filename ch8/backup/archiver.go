package backup

type Archiver interface {
	Archiver(src, dest string) error
}

type zipper struct {
}
