from locust import HttpUser, SequentialTaskSet, task, between
import faker, json, os

BASE_URL = "/otusapp/api/v1"

fake = faker.Faker()
headers = {"Host": os.environ["SERVICE_HOST"]}

class UserBehavior(SequentialTaskSet):
    def __init__(self, parent):
       super(UserBehavior, self).__init__(parent)
       self.user_id = ""

    @task
    def create_user_task(self):
        first_name = fake.first_name()
        last_name = fake.last_name()
        username = first_name.lower() + "_" + last_name.lower()

        response = self.client.post(BASE_URL + "/users", data = json.dumps({
                "username": username,
                "firstName": first_name,
                "lastName": last_name,
                "email": fake.email(),
                "phone": fake.phone_number()
		    }),
            name = "CreateUser",
            headers = headers,
            catch_response = True)

        if response.status_code == 200:
            try:
                self.user_id = response.json().get("id")
                response.success()
            except ValueError:
                response.failure("failed to parse user_id from response body")
                self.interrupt()
        else:
            response.failure(f'status code is {response.status_code}')
            self.interrupt()

    @task
    def find_user_by_id_task(self):
        if self.user_id == "":
            return

        self.client.get(BASE_URL + "/users/" + self.user_id, name = "FindUserById", headers = headers)

    @task
    def update_user_task(self):
        if self.user_id == "":
            return

        self.client.put(BASE_URL + "/users/" + self.user_id, data = json.dumps({
                "firstName": fake.first_name(),
                "lastName": fake.last_name(),
                "email": fake.email(),
                "phone": fake.phone_number()
		    }),
            name = "UpdateUser",
            headers = headers)

    @task
    def list_users_task(self):
        self.client.get(BASE_URL + "/users", name = "ListUsers", headers = headers)

    @task
    def delete_user_task(self):
        if self.user_id == "":
            return

        self.client.delete(BASE_URL + "/users/" + self.user_id, name = "DeleteUser", headers = headers)

class ApiUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(3, 5)