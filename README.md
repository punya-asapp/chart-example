# chart-example

Demonstrates a CI workflow for Docker images and Helm charts, with the following properties:
* A new Docker image is generated on each commit to `master`, with a unique tag that ties back to the relevant source commit.
* If the source commit coincides with a Git tag, the Docker image's tag is human readable.
* Each chart refers to a *specific* Docker image version.
* Changes that only affect configuration (i.e. are confined to the `helm-charts` directory) do not trigger an image build, and the Helm chart refers to a pre-existing image built from the last commit that touched sources.

## How it works

This repository uses:

* Github actions to run CI scripts
* Github packages to store Docker images
* Github artifacts to store generated charts

An alternative implementation might use CircleCI to run CI scripts, Amazon ECS to store Docker images, and S3 to store charts. The core idea would remain the same.

The following steps are taken:

1. Each commit triggers a CI run.

2. In a CI run, find the commit ID of the most recent commit that touched a source file ([details]).

3. Check to see if a Docker image derived from that commit has been previously published. If not, build, tag and publish it.

4. Generate and publish a chart by putting this Docker image name into `values.yaml`.

[details]: https://stackoverflow.com/questions/5685007/making-git-log-ignore-changes-for-certain-paths
