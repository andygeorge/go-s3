import (
	"context"
	"fmt"

	aws_s3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Reader struct {
	ctx        context.Context
	bucketName string
	client     *aws_s3.Client
}

func main() {
	ctx := context.TODO()
	s3Reader, err := getS3Reader(ctx)
	if err != nil {
		// handle err
	}

	s3Reader.bucketName = "ur bucket"
	targetFolder = "ur folder in bucket"

	s3Keys, err = allS3KeysInFolder(s3Reader, target)
	if err != nil {
		// handle err
	}

	// print all files in folder
	for _, s3Key := range s3Keys {
		fmt.Print(s3Key)
	}
}

func getS3Reader(ctx context.Context) (S3Reader, error) {
	awsConfig, err := aws_config.LoadDefaultConfig(ctx)
	if err != nil {
		return S3Reader{ctx, "", nil}, err
	}

	return S3Reader{
		ctx:        ctx,
		bucketName: "",
		client:     aws_s3.NewFromConfig(awsConfig),
	}, nil
}

func allS3KeysInFolder(s3Reader S3Reader, folder string) ([]string, error) {
	var s3keys []string
	objects := &aws_s3.ListObjectsV2Input{
		Bucket: aws.String(s3Reader.bucketName),
		Prefix: aws.String(folder),
	}

	paginator := aws_s3.NewListObjectsV2Paginator(s3Reader.client, objects)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}

		for _, obj := range output.Contents {
			s3keys = append(s3keys, *obj.Key)
		}
	}

	return s3keys, nil
}
