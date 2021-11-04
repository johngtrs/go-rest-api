docker compose down

volumes=$(docker volume ls -q)

if [[ -n "${volumes// /}" ]]; then
    docker volume rm $(docker volume ls -q)
fi

if [ $# -eq 0 ]; then
    echo "no images deleted"
elif [ $1 = "--all" ]; then
    docker rmi -f $(docker images -a -q)
elif [ $1 = "--web" ]; then
    # Delete web image with format foldername_web (default docker naming)
    image="${PWD##*/}_web"
    docker rmi -f $(docker images -a -q $image)
fi