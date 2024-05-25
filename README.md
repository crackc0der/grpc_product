RUN: make up

**Categoryis**:
    GETALL:
        GET: /categories

    ADD:
        {
            "CategoryName: "category_name"
        }

        POST: /category

    UPDATE:
        {
            "Id": category_id
        }

        PUT: /category

    DELETE:
        {
            "Id": category_id
        }

        DELETE: /category

**Peoducts**:
    GETALL:
        GET: /products

    GET:
        {
            "Id": product_id
        }

        GET: /product

    ADD:
        {
            "ProductName": "product_name",
            "ProductCategory": product_category_id,
            "ProductPrice": peoduct_price
        }

        POST: /product

    UPDATE:
        {
            "Id": product_id,
            "ProductName": "product_name",
            "ProductCategory": product_category_id,
            "ProductPrice": peoduct_price            
        }

        PUT: /product

    DELETE:
        {
            "ProductId": product_id
        }

        DELETE: /product