# terraform-s3-dir

``terraform-s3-dir`` is a small utility that takes a directory of files and produces a configuration file for [Terraform](https://terraform.io/) that will upload those files into a particular named S3 bucket.

It could be useful for using Terraform to deploy a static website to S3's website publishing feature.

This utility just generates the ``aws_s3_bucket_object`` configurations. It's up to the user to separately create the bucket into which the objects will be placed.
