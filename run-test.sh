GATLING_BIN_DIR=${PWD}/bin

WORKSPACE=${PWD}

sh $GATLING_BIN_DIR/gatling.sh -rm local -s NomeDaSimulacao

sleep 1

curl -v "http://localhost:3000/health"
