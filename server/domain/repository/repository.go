package repository

import "github.com/opensourceways/sync-repository-file/server/domain"

type RepoFile interface {
	SaveFile(domain.PlatformOrgRepo, string, domain.RepoFile) error
	FindFiles(repo domain.PlatformOrgRepo, branch, fileName string) ([]domain.RepoFileInfo, error)
}
