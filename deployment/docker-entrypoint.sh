#!/usr/bin/env sh

cp /invokes.yml.template /conf_dir/mh_conf.yml.sample

if [[ "$DB_ENGINE" == "mysql" ]] ; then
    export -p DB_CONN_STRING="$DB_PROTO$DB_USER:$DB_PASSWD@tcp($DB_HOST:$DB_PORT)/$DB_NAME$DB_OPTS"
fi

if [[ "$DB_ENGINE" == "postgresql" ]] ; then
    export -p DB_CONN_STRING="$DB_PROTO$DB_USER:$DB_PASSWD@$DB_HOST:$DB_PORT"
fi

if [ ! -f /conf_dir/mh_conf.yml ] ; then
    cat /invokes.yml.template |
        sed "s|<engine>|$DB_ENGINE|g" | 
        sed "s|<conn_string>|$DB_CONN_STRING|g" | 
        sed "s|<cors_allowed_origins>|$CORS_ALLOWED_ORIGINS|g" |  
        sed "s|<cors_methods>|$CORS_ALLOWED_METHODS|g" |  
        sed "s|<cors_headers>|$CORS_ALLOWED_HEADERS|g" |  
        sed "s|<cors_max_age>|$CORS_MAX_AGE|g" |  
        sed "s|<log_consolelevel>|$LOG_LEVEL|g" | 
        sed "s|<log_usefile>|false|g" | 
        sed "s|<log_filelevel>|debug|g" | 
        sed "s|<log_filepath>|.|g" | 
        sed "s|<log_filemaxsize>|50|g" | 
        sed "s|<log_filemaxbackup>|3|g" | 
        sed "s|<log_filemaxage>|28|g" | 
        sed "s|<port>|$PORT|g" |
        sed "s|<basepath>|$PROJET_BASE_PATH|g" > /conf_dir/mh_conf.yml
fi

if [[ "$CAT_CONF" == "true" ]] ; then
    cat /conf_dir/mh_conf.yml
fi

echo "Waiting for database to be up..."
while ! nc -z ${DB_HOST} ${DB_PORT}; do sleep 1; done
echo "Connected to database."

exec "$@"