import requests

response = requests.get('http://localhost:3000/employees')
if response.status_code == 200:
    employees = response.json()
    print(employees)
else:
    print(f'Failed to retrieve posts: {response.status_code}')