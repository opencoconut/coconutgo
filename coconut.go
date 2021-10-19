package coconut

type Client struct {
	APIKey       string
	Region       string
	Endpoint     string
	Storage      Storage
	Notification Notification

	Job      *JobClient
	Metadata *MetadataClient
}

type Storage map[string]interface{}
type StorageCredentials map[string]string
type Notification map[string]interface{}

type JobClient struct {
	Client Client
}

type MetadataClient struct {
	Client Client
}

const ENDPOINT = "https://api.coconut.co/v2"

func NewClient(cli Client) Client {
	cli.Job = &JobClient{Client: cli}
	cli.Metadata = &MetadataClient{Client: cli}

	return cli
}

func (c Client) GetEndpoint() string {
	if c.Endpoint != "" {
		return c.Endpoint
	}

	if c.Region != "" {
		return "https://api-" + c.Region + ".coconut.co/v2"
	}

	return ENDPOINT
}
