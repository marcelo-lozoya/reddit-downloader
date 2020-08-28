package downloader

import (
	"testing"

	"github.com/marcelo-lozoya/reddit-downloader/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type downloaderMockFields struct {
	DirectoryHelper *mocks.IDirectoryHelper
}

// Test for the Downloader.MakeRequestForReddit function
func TestCDownloader_MakeRequestForReddit(t *testing.T) {
	type args struct {
		subreddit string
		limit     int
	}

	tests := []struct {
		name   string
		fields downloaderMockFields
		args   args
	}{

		// Test cases go here
		{
			name:   "test MakeRequestForReddit",
			fields: downloaderMockFields{DirectoryHelper: &mocks.IDirectoryHelper{}},
			args:   args{subreddit: "aww", limit: 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Here we can tell the Downloader's dependencies how they should behave.
			// In this case, we want createRequiredFolders to do nothing, causing MakeRequestForReddit
			// to panic because it will try to write files to a directory that doesn't exist
			tt.fields.DirectoryHelper.On("createRequiredFolders", mock.Anything).Return(nil)

			// Instantiate a Downloader with our mocked DirectoryHelper
			d := CDownloader{DirectoryHelper: tt.fields.DirectoryHelper}

			// Expect a panic
			assert.Panics(t, func() {
				d.MakeRequestForReddit(tt.args.subreddit, tt.args.limit)
			})
		})
	}
}
