/*
* Create DB command could be used if necessary:
* */

CREATE DATABASE wallester;

CREATE TABLE public.customers
(
    id           uuid         NOT NULL,
    first_name   varchar(100) NOT NULL,
    last_name    varchar(100) NOT NULL,
    birthdate    timestamp    NOT NULL,
    email        varchar(255) NOT NULL,
    gender       varchar(25)  NOT NULL,
    search_field varchar(210) NOT NULL,
    created_at   timestamp    NOT NULL,
    updated_at   timestamp    NOT NULL,
    CONSTRAINT customers_pkey PRIMARY KEY (id)
);

-- Permissions
ALTER TABLE public.customers OWNER TO postgres;
GRANT
ALL
ON TABLE public.customers TO postgres;

/*
 * Depends on business requirement should be added indexes.
 * E.g. if a strict search will be required very often for the "first_name" and "last_name" fields,
 * it makes sense to create an index for the first_name and last_name field, etc...
 * */

INSERT INTO customers
(id,first_name,last_name,birthdate,email,gender,search_field,created_at,updated_at)
VALUES
    ('b737ec90-c69d-493e-8e78-4bd26a686a72','Charles','Cameron','2003-10-22 00:00:00.000','CharlesCameron@test.com','Male','charles cameron','2021-10-24 14:07:18.124','2021-10-24 14:07:18.124'),
    ('c1d4fe8f-b564-4b9f-a235-c2b73da43016','David','Cameron','2003-09-09 00:00:00.000','DavidCameron@test.test','Male','david cameron','2021-10-24 14:07:52.032','2021-10-24 14:07:52.032'),
    ('dc1b69b6-a436-4fa9-9479-2f0c78ff6edf','Gervase','Cameron','2003-08-13 00:00:00.000','GervaseCameron@test.test','Male','gervase cameron','2021-10-24 14:08:17.321','2021-10-24 14:08:17.321'),
    ('255dd40c-27f8-413f-ab34-846b919803c0','Lesley','Cameron','2003-07-15 00:00:00.000','LesleyCameron@test.test','Female','lesley cameron','2021-10-24 14:08:39.817','2021-10-24 14:08:58.578'),
    ('392e2380-9b20-4737-b298-a79018ca5c64','Lesley','Ross','2002-11-13 00:00:00.000','LesleyRoss@test.test','Female','lesley ross','2021-10-24 14:09:30.434','2021-10-24 14:09:30.434'),
    ('f62c65a7-e655-4ee7-829c-2f144581baab','Jane','Ross','1964-07-14 00:00:00.000','JaneRoss@test.test','Female','jane ross','2021-10-24 14:10:16.311','2021-10-24 14:10:16.311'),
    ('b33836bf-75c7-4f7f-93ab-1c164fa62d43','Mary','Brooks','1965-05-10 00:00:00.000','MaryBrooks@test.test','Female','mary brooks','2021-10-24 14:11:43.938','2021-10-24 14:11:43.938'),
    ('5b013820-af15-4cac-a8de-ed7b45a6be8a','Joan','Brooks','1998-12-20 00:00:00.000','JoanBrooks@test.test','Female','joan brooks','2021-10-24 14:12:13.474','2021-10-24 14:12:13.474'),
    ('fe7e91c6-da63-4bab-a537-2edd18f7fced','John','Brooks','1984-02-17 00:00:00.000','JohnBrooks@test.test','Male','john brooks','2021-10-24 14:12:52.216','2021-10-24 14:12:52.216'),
    ('0913ea65-399b-4a3e-8848-5676af5b2298','Morris','Stephens','2000-07-13 00:00:00.000','MorrisStephens@test.com','Male','morris stephens','2021-10-24 14:13:33.213','2021-10-24 14:13:33.213'),
    ('6182b177-38cc-4ad4-bf00-53a4ea8de98f','Paul','Stephens','1996-05-18 00:00:00.000','PaulStephens@test.com','Male','paul stephens','2021-10-24 14:14:07.695','2021-10-24 14:14:07.695'),
    ('3f8a5a51-703b-4241-b578-22d5682299b5','Charlotte','Stephens','1992-04-29 00:00:00.000','CharlotteStephens@test.com','Female','charlotte stephens','2021-10-24 14:14:42.267','2021-10-24 14:14:42.267'),
    ('b7de6b37-492f-4f39-b818-9fa6a37ce668','Ruth','Stephens','1989-01-20 00:00:00.000','RuthStephens@test.com','Female','ruth stephens','2021-10-24 14:15:14.133','2021-10-24 14:15:14.133'),
    ('db79903a-a1c8-4420-98ee-fee45e05ae1c','Ilene','Evans','2003-10-06 00:00:00.000','IleneEvans@test.com','Female','ilene evans','2021-10-24 14:15:41.919','2021-10-24 14:15:41.919'),
    ('38f60a00-a857-4cab-8333-cf2e6fc00357','Charleen','Evans','2002-09-02 00:00:00.000','CharleenEvans@test.com','Female','charleen evans','2021-10-24 14:16:13.084','2021-10-24 14:16:13.084'),
    ('019f657e-4ffd-4dc2-86cf-58b5c58b08dc','William','Ward','2003-05-11 00:00:00.000','WilliamWard@test.com','Male','william ward','2021-10-24 14:16:49.638','2021-10-24 14:16:49.638'),
    ('556308c7-bb35-4b03-8547-9b1791ac8dc2','William','Lawson','1986-07-29 04:00:00.000','WilliamLawson@test.com','Male','william lawson','2021-10-24 14:17:19.642','2021-10-24 14:17:19.642');
