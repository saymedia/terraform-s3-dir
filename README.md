# terraform-s3-dir

``terraform-s3-dir`` is a small utility that takes a directory of files and produces a configuration file for [Terraform](https://terraform.io/) that will upload those files into a particular named S3 bucket.

It could be useful for using Terraform to deploy a static website to S3's website publishing feature.

This utility just generates the ``aws_s3_bucket_object`` configurations. It's up to the user to separately create the bucket into which the objects will be placed.

## Installing

Pretty standard Go program:

* ``go get github.com/saymedia/terraform-s3-dir``
* ``go install github.com/saymedia/terraform-s3-dir``

## Usage

```
Usage: terraform-s3-dir [-h] [-x glob patterns to exclude] <root dir> <bucket name>
 -h, --help
 -x, --exclude=glob patterns to exclude
```

