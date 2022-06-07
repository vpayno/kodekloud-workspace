#!/bin/bash

echo -e " Log name   \t      GET      \t      POST    \t   DELETE "
echo -e "------------------------------------------------------------"

mapfile -t app_names < ./resources/apps.txt

for app_name in "${app_names[@]}"; do
	get_requests=$(grep -c "GET" < "./resources/logs/${app_name}"_app.log)
	post_requests=$(grep -c "POST" < "./resources/logs/${app_name}"_app.log)
	delete_requests=$(grep -c "DELETE" < "./resources/logs/${app_name}"_app.log)

	#echo -e " Finance    \t ${get_requests}    \t    ${post_requests}   \t   ${delete_requests}"
	printf " %-10s \t %d    \t    %d   \t   %d\n" "${app_name}" "${get_requests}" "${post_requests}" "${delete_requests}"
done
