sh register_test.sh
sh offer_post.sh
sh offer_post.sh

curl --location --request GET 'http://127.0.0.1:9999/offer/offer_list'
echo '\n'

curl --location --request GET 'http://127.0.0.1:9999/offer/offer_list/user?uid=1'
echo '\n'
curl --location --request GET 'http://127.0.0.1:9999/offer/offer_list/user?uid=2'
echo '\n'
curl --location --request GET 'http://127.0.0.1:9999/offer/offer_list/user?uid=3'
echo '\n'
curl --location --request GET 'http://127.0.0.1:9999/offer/offer_list/user?uid=666'
echo '\n'