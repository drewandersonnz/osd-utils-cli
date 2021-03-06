## osdctl account clean-velero-snapshots

Cleans up S3 buckets whose name start with managed-velero

### Synopsis

Cleans up S3 buckets whose name start with managed-velero

```
osdctl account clean-velero-snapshots [flags]
```

### Options

```
  -a, --aws-access-key-id string       AWS Access Key ID
  -c, --aws-config string              specify AWS config file path
  -p, --aws-profile string             specify AWS profile
  -r, --aws-region string              specify AWS region (default "us-east-1")
  -x, --aws-secret-access-key string   AWS Secret Access Key
  -h, --help                           help for clean-velero-snapshots
```

### Options inherited from parent commands

```
      --cluster string             The name of the kubeconfig cluster to use
      --context string             The name of the kubeconfig context to use
      --insecure-skip-tls-verify   If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string          Path to the kubeconfig file to use for CLI requests.
      --request-timeout string     The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string              The address and port of the Kubernetes API server
```

### SEE ALSO

* [osdctl account](osdctl_account.md)	 - AWS Account related utilities

