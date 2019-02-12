package filetypes

import (
	"fmt"

	"github.com/juju/errors"
)

// isSameVideos checks if two videos (local and uploaded) are exactly the same
func isSameVideos(uploadedFileURL, localFilePath string) bool {
	upHash, err := fileHash(uploadedFileURL)
	if err != nil {
		return false
	}
	localHash, err := fileHash(localFilePath)
	if err != nil {
		return false
	}

	return upHash == localHash
}

// VideoTypedMedia implements TypedMedia for video files
type VideoTypedMedia struct{}

// IsCorrectlyUploaded checks that the video that was uploaded is the same as the local one, before deleting the local one
func (gm *VideoTypedMedia) IsCorrectlyUploaded(uploadedFileURL, localFilePath string) (bool, error) {
	if !IsVideo(localFilePath) {
		return false, fmt.Errorf("%s is not a video. Not deleting local file", localFilePath)
	}

	// compare uploaded image and local one
	if isSameVideos(uploadedFileURL, localFilePath) {
		return true, nil
	}

	return false, errors.Errorf("Not sure if video was uploaded correctly. Not deleting local file. URL: %s", uploadedFileURL)
}
