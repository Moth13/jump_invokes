#!/usr/bin/env sh

cp /app.yml.template /conf_dir/mh_conf.yml.sample

if [ ! -f /conf_dir/mh_conf.yml ] ; then
    cat /app.yml.template |
        sed "s|<cors_allowed_origins>|$CORS_ALLOWED_ORIGINS|g" |  
        sed "s|<cors_methods>|$CORS_ALLOWED_METHODS|g" |  
        sed "s|<cors_headers>|$CORS_ALLOWED_HEADERS|g" |  
        sed "s|<cors_max_age>|$CORS_MAX_AGE|g" |  
        sed "s|<debug>|$DEBUG|g" | 
        sed "s|<port>|$PORT|g" |
        sed "s|<origin>|$FS_ORIGIN|g" |
        sed "s|<basepath>|$PROJET_BASE_PATH|g" |
        sed "s|<installpath>|/root/go/src/app|g" |
        sed "s|<version_prefixed_routes>|$VERSION_PREFIXED_ROUTES|g" > /conf_dir/mh_conf.yml
fi

cat /root/go/src/app/view/html/404.html |
    sed "s|/css/404.css|/app/css/404.css|g" > /root/go/src/app/view/html/404.html

if [[ "$CAT_CONF" == "true" ]] ; then
    cat /conf_dir/mh_conf.yml
fi

exec "$@"