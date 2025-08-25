# API Contract Testing with Postman

## Example: GET /api/users

Add this to the "Tests" tab in your Postman request:

```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

const userSchema = {
    "type": "object",
    "properties": {
        "id": { "type": "string" },
        "email": { "type": "string", "format": "email" },
        "role": { "type": "string" },
        "status": { "type": "string" }
    },
    "required": ["id", "email", "role", "status"]
};

pm.test("Each user matches schema", function () {
    const users = pm.response.json();
    users.forEach(user => {
        pm.expect(tv4.validate(user, userSchema), JSON.stringify(tv4.error)).to.be.true;
    });
});
```

---

## Example: POST /api/auth/login

Add this to the "Tests" tab in your Postman request:

```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

const loginSchema = {
    "type": "object",
    "properties": {
        "token": { "type": "string" }
    },
    "required": ["token"]
};

pm.test("Login response matches schema", function () {
    pm.response.to.have.jsonSchema(loginSchema);
});
```

---

## How to Use
- Paste these scripts into the "Tests" tab of your Postman request.
- Adjust the schema to match your actual API response.
- Use `pm.response.to.have.jsonSchema(schema)` for single objects, or loop for arrays.
- For CI/CD, export your collection and run with Newman.
