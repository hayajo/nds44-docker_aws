#!/bin/sh

MACHINE_DRIVER=${MACHINE_DRIVER:-"virtualbox"}
MACHINE_OPTS=${MACHINE_OPTS:-""}
MACHINE_NAME="dev"

if [ "$MACHINE_DRIVER" == "amazonec2" ]; then
  MACHINE_NAME="$MACHINE_NAME-$MACHINE_DRIVER"
  MACHINE_OPTS="$MACHINE_OPTS \
  --amazonec2-access-key=$AWS_ACCESS_KEY_ID \
  --amazonec2-secret-key=$AWS_SECRET_ACCESS_KEY \
  --amazonec2-ami=$AWS_AMI \
  --amazonec2-region=$AWS_DEFAULT_REGION \
  --amazonec2-vpc-id=$AWS_VPC_ID"
fi

set -x
docker-machine create \
  --driver=$MACHINE_DRIVER \
  $MACHINE_OPTS \
  $MACHINE_NAME
set +x

echo "run: eval \"\$(docker-machine env $MACHINE_NAME)\""
