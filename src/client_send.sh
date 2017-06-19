curl -d "{\"id\":\"01232\",\"classNumber\":1,\"score\":97}" localhost:5000/register-student
curl -d "{\"id\":\"01233\",\"classNumber\":1,\"score\":98}" localhost:5000/register-student
curl -d "{\"id\":\"01234\",\"classNumber\":1,\"score\":99}" localhost:5000/register-student
curl -d "{\"classNumber\":1,\"teacher\":\"teacher1\"}" localhost:5000/register-student
curl localhost:5000/get-class-total-score/12345