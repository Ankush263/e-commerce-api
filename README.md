# Model Structure
- ## E - commerce API
    - ### User Model
        - id 
        - username
        - email
        - password
        - phone
        - role 
            - enum
                - seller
                - customer
        - created_at
        - updated_at
    - ### Store Model
        - id
        - user
        - name
        - description
        - store_type
        - store_id
        - created_at
        - updated_at
    - ### Product Model
        - id
        - name
        - description
        - owner
        - store_id
        - price
        - created_at
        - updated_at