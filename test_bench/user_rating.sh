curl --location --request POST 'http://127.0.0.1:9999/user_rating/1' \
--form 'offer_id="1"' \
--form 'rating="10"' \
--form 'comment="good"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/1' \
--form 'offer_id="2"' \
--form 'rating="5"' \
--form 'comment="normal"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/2' \
--form 'offer_id="3"' \
--form 'rating="7"' \
--form 'comment="good"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/3' \
--form 'offer_id="4"' \
--form 'rating="3"' \
--form 'comment="bad"'

echo '\n'

curl --location --request POST 'http://127.0.0.1:9999/user_rating/1' \
--form 'offer_id="5"' \
--form 'rating="7"' \
--form 'comment="good"'

echo '\n'

# 不正确的offer_id
curl --location --request POST 'http://127.0.0.1:9999/user_rating/2' \
--form 'offer_id="100"' \
--form 'rating="5"' \
--form 'comment="normal"'

echo '\n'