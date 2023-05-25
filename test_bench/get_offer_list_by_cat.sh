sh register_test.sh
sh offer_post.sh
sh offer_post.sh
#获取'拿外卖'分类下的offers
curl --location --request GET 'http://127.0.0.1:9999/offer_list_by_cat/拿外卖'
echo '\n'

#获取'跑腿'分类下的offers
curl --location --request GET 'http://127.0.0.1:9999/offer_list_by_cat/跑腿'
echo '\n'

#获取不存在的分类下的offers 应返回空列表或者错误
curl --location --request GET 'http://127.0.0.1:9999/offer_list_by_cat/non_existent_category'
echo '\n'
