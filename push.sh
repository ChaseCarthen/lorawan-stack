# pushing to ncar-da-15 $1 is image name, $2 is the tag we want it to be on kubernetes, $3 is latest
dockertag=ncar-da-15.rc.unr.edu/library/$1:$3
echo $dockertag
docker tag $1:$2 $dockertag
docker push $dockertag
