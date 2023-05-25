#!/bin/bash
sh register_test.sh

test_update_user_info() {
    USER_ID=$1
    NICKNAME=$2
    FULLNAME=$3
    PHONE_NUMBER=$4

    JSON_DATA=$(cat <<EOF
{
	"NickName": "$NICKNAME",
	"FullName": "$FULLNAME",
	"PhoneNumber": "$PHONE_NUMBER"
}
EOF
)

    echo "Running test: updating user $USER_ID info"
    curl -X PUT -H "Content-Type: application/json" -d "$JSON_DATA" http://127.0.0.1:9999/change_usr_info/$USER_ID
    echo ""
}

# Test with a valid user ID and valid user info
test_update_user_info 1 "new_nickname1" "new_fullname1" "new_phone_number1"

# Test with a non-existent user ID
test_update_user_info 999999 "new_nickname2" "new_fullname2" "new_phone_number2"

# Test with an empty user ID
test_update_user_info "" "new_nickname3" "new_fullname3" "new_phone_number3"

# Test with a valid user ID but empty user info
test_update_user_info 1 "" "" ""

# Test with a valid user ID and partially empty user info
test_update_user_info 1 "new_nickname4" "" ""
