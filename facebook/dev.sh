export FACEBOOK_DB_HOST="localhost:5432"
export FACEBOOK_DB_NAME="fb"
export FACEBOOK_DB_USER="spud"
export FACEBOOK_DB_PASSWORD="stacey"

docker run --name "facebook-db" \
  -e POSTGRES_DB="$FACEBOOK_DB_NAME" \
  -e POSTGRES_USER="$FACEBOOK_DB_USER" \
  -e POSTGRES_PASSWORD="$FACEBOOK_DB_PASSWORD" \
  -p 5432:5432 \
  -d postgres
