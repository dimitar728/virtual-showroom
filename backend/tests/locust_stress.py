from locust import HttpUser, task, between

class WebsiteUser(HttpUser):
    wait_time = between(1, 3)

    @task
    def login_and_book(self):
        login = self.client.post("/api/auth/login", json={
            "email": "user@example.com",
            "password": "password123"
        })
        if login.status_code == 200:
            token = login.json().get("token")
            headers = {"Authorization": f"Bearer {token}"}
            self.client.get("/api/showrooms", headers=headers)
            self.client.post("/api/bookings", json={
                "showroomId": "1",
                "slotTime": "2025-08-26T10:00:00Z"
            }, headers=headers)
