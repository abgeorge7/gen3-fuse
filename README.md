# gen3-fuse


gen3-fuse is an S3 file system built on Goofys that uses the workspace-token-service. 

----
## Overview

This repository supports the same functionality as Goofys, but doesn't require user to have AWS credentials. Instead, we use the [workspace-token-service](https://github.com/uc-cdis/workspace-token-service) for authentication.

See the [goofys documentation](https://github.com/kahing/goofys) for a discussion of FUSE and s3. 

----

## Setup

----

Use the --recursive flag when you clone this repository to ensure you download the content of any submodules.

    git clone --recursive https://github.com/uc-cdis/gen3-fuse.git
    git submodule update --init --recursive

----
## References
* [goofys](https://github.com/kahing/goofys)
