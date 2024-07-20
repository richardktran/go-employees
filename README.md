## How to run the code
1. Clone the repository
```bash
git clone https://github.com/richardktran/go-employees.git
cd go-employees
```
2. Fetch dependencies
```bash
go mod download
```
3. Sync data from https://dummy.restapiexample.com/api/v1/employees to txt file
```bash
make sync-data
```
4. Run the server
```bash
make start
```
5. To see the list of employees, run the curl command below or visit http://localhost:8080/employees in your browser
```bash
curl http://localhost:8080/employees
```
6. To add new employee, run the curl command below
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Richard Tran","salary":100000,"age":24, "profile_image":"profile1.png"}' http://localhost:8080/employees
```