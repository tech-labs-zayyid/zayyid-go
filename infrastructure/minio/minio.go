package minio

// func MinioConnection(config config.EnvironmentConfig) (*minioMain.Client, error) {
// 	endpoint := config.StorageMinioServer
// 	accessKeyID := config.StorageMinioAccessKey
// 	secretAccessKey := config.StorageMinioSecreatKey
// 	suseSSL := config.StorageMinioUseSSL
// 	useSSL, _ := strconv.ParseBool(suseSSL)

// 	// Initialize minio client object.
// 	minioClient, errInit := minioMain.New(endpoint, &minioMain.Options{
// 		Creds:  minioCredentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if errInit != nil {
// 		log.Fatalln(errInit)
// 	}

// 	return minioClient, errInit
// }
