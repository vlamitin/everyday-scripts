BUILD_DATE=`date +%d_%m_%YT%H_%M_%S`

OKKO_FILMS_BINARY=okko_films

OPTIMUM_COLLECTION=msvod_all_optimum
NEW_PROMO_COLLECTION=new-promo

WINDOWS_ENV=env GOOS=windows GOARCH=amd64
LINUX_ENV=env GOOS=linux GOARCH=amd64

OKKO_OPTIMUM_FLAGS=-ldflags "-w -s -X main.Collection=${OPTIMUM_COLLECTION} -X main.BuildDate=${BUILD_DATE}"
OKKO_NEW_PROMO_FLAGS=-ldflags "-w -s -X main.Collection=${NEW_PROMO_COLLECTION} -X main.BuildDate=${BUILD_DATE}"

# Builds okko_films binaries for amd64 linux and amd64 windows
build_okko_films:
	${WINDOWS_ENV} go build ${OKKO_OPTIMUM_FLAGS} -o bin/${OKKO_FILMS_BINARY}_${OPTIMUM_COLLECTION}_${BUILD_DATE}.exe cmd/okko_films_to_csv/films_to_csv.go
	${WINDOWS_ENV} go build ${OKKO_NEW_PROMO_FLAGS} -o bin/${OKKO_FILMS_BINARY}_${NEW_PROMO_COLLECTION}_${BUILD_DATE}.exe cmd/okko_films_to_csv/films_to_csv.go
	${LINUX_ENV} go build ${OKKO_OPTIMUM_FLAGS} -o bin/${OKKO_FILMS_BINARY}_${OPTIMUM_COLLECTION}_${BUILD_DATE} cmd/okko_films_to_csv/films_to_csv.go
	${LINUX_ENV} go build ${OKKO_NEW_PROMO_FLAGS} -o bin/${OKKO_FILMS_BINARY}_${NEW_PROMO_COLLECTION}_${BUILD_DATE} cmd/okko_films_to_csv/films_to_csv.go

# Cleans our project: deletes binaries
clean_okko_films:
	rm -f bin/${OKKO_FILMS_BINARY}*
