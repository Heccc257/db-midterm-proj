curl --location --request POST 'http://127.0.0.1:9999/user_rating/1' \
--form 'offer_id="1"' \
--form 'rating="10"' \
--form 'comment="good"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/2' \
--form 'offer_id="3"' \
--form 'rating="5"' \
--form 'comment="normal"'

echo '\n'

# 不正确的offer_id
curl --location --request POST 'http://127.0.0.1:9999/user_rating/2' \
--form 'offer_id="100"' \
--form 'rating="5"' \
--form 'comment="normal"'

echo '\n'