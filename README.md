# lfs-livy
 A livy server to be used to run, monitor and cancel spark jobs for the LFS Survey

##Getting Started With Livy

Download apache-livy

`https://livy.apache.org/download/`

####Config:

In the  `livy.conf.template` file, set the spark master and deploy modes:

`livy.spark.master = local[*]`

`livy.spark.deploy-mode = client`

Add local file directories to the whitelist (any directories containing any local files being used)

`livy.file.local-dir-whitelist = /absolute/path/to/directory`

Rename file to `livy.conf`

####Environment

In the `livy-env.sh.template` file, add Spark_Home, below is path for a brew installation of spark:

`export SPARK_HOME="/usr/local/Cellar/apache-spark/2.4.5/libexec"`

Any environment variables can also be added in the same way.

Rename file to `livy-env.sh`

####Starting server

From the livy root folder, the livy server can be started and stopped with:

`./bin/livy-server start`

`./bin/livy-server stop`

Livy can be viewed from the livy UI at: `http://localhost:8998/ui/`

###Uploading Jars

In the application going into livy a fat jar can be created using:

`sbt assembly`

This will encompass all the external dependencies.

The absolute path of this jar goes under 'file' in the request and the folder containing it must be covered by the whitelist.
