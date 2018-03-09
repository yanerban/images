package models

import (
	"fmt"
	"github.com/weilaihui/fdfs_client"
)

func FDFSUploadfilename(filename string) (groupName string, fielId string, err error) {
	fdfs_upclient, err := fdfs_client.NewFdfsClient("./conf/client.conf")
	if err != nil {
		fmt.Println("New Fdfsclient err = ", err)
		return "", "", err
	}

	uploadResponse, err := fdfs_upclient.UploadByFilename(filename)
	if err != nil {
		fmt.Println("uploadbyfilename err = ", err)
		return "", "", err
	}

	fmt.Println(uploadResponse.GroupName)
	fmt.Println(uploadResponse.RemoteFileId)

	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil
}

func FDFSUploadByBuffer(buffer []byte, suffix string) (groupName string, fielId string, err error) {
	fdfs_upclient, err := fdfs_client.NewFdfsClient("./conf/client.conf")
	if err != nil {
		fmt.Println("New FDFSclient err = ", err)
		return "", "", err
	}

	uploadResponse, err := fdfs_upclient.UploadByBuffer(buffer, suffix)
	if err != nil {
		fmt.Println("uploadbybuffer err = ", err)
		return "", "", err
	}

	fmt.Println(uploadResponse.GroupName)
	fmt.Println(uploadResponse.RemoteFileId)

	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil

}

func FDFSDownloadToFile(filename string) (size int64, fielId string, err error) {
	fdfs_downclient, err := fdfs_client.NewFdfsClient("./conf/client.conf")
	if err != nil {
		fmt.Println("New FDFSclient err = ", err)
		return 0, "", err
	}

	uploadResponse, err := fdfs_downclient.UploadByFilename(filename)
	defer fdfs_downclient.DeleteFile(uploadResponse.RemoteFileId)
	if err != nil {
		fmt.Println("uploadfilename err = ", err)
		return 0, "", err
	}

	downloadResponse, err := fdfs_downclient.DownloadToFile(filename, uploadResponse.RemoteFileId, 0, 0)
	if err != nil {
		fmt.Println("dowloadfilename err = ", err)
		return 0, "", err
	}

	fmt.Println(downloadResponse.DownloadSize)
	fmt.Println(downloadResponse.RemoteFileId)

	return downloadResponse.DownloadSize, downloadResponse.RemoteFileId, nil

}

func FDFSDownloadToBuffer(fielid string, downloadsize int64) (size int64, fielId string, err error) {
	fdfs_downclient, err := fdfs_client.NewFdfsClient("./conf/client.conf")
	if err != nil {
		fmt.Println("New FDFSclient err = ", err)
		return 0, "", err

	}
	downloadResponse, err := fdfs_downclient.DownloadToBuffer(fielid, 0, downloadsize)
	if err != nil {
		fmt.Println("downloadfile err = ", err)
		return 0, "", err
	}

	fmt.Println(downloadResponse.DownloadSize)
	fmt.Println(downloadResponse.RemoteFileId)

	return downloadResponse.DownloadSize, downloadResponse.RemoteFileId, nil
}
