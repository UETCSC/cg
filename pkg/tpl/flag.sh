#!/bin/bash
<<PROMPT
Please modify this script as needed to ensure that the flag can be placed in the correct location
Some web questions require correct domain name access. Please configure the correct domain name in this script.

FLAG The correct flag passed in by the platform
DOMAIN The domain name of the current environment passed in by the platform

The following are the built-in functions

Write the flag to the file system
#Write /flag by default
write_flag_in_fs

# Write to the web directory /var/www/html/flag.txt
write_flag_in_fs /var/www/html/flag.txt
PROMPT

write_flag_in_fs() {
    # 将flag写入到文件系统中
    if [ -z "$1" ]; then
        flag_path="/flag"
    else
        flag_path="$1"
    fi
    echo ${FLAG} > ${flag_path}
}

write_flag_in_db() {
    local db_name="${1:-web}"
    local db_table="${2:-flag}"
    local db_column="${3:-flag}"
    echo mysql -uroot -proot -e "update ${db_name}.${db_table} set ${db_column}='${FLAG}';"
}

export FLAG=not_flag
FLAG=not_flag