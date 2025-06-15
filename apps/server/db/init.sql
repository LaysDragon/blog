INSERT INTO public.account
( deleted_at, username, "role", email, passwd_hash)
VALUES(NULL, 'admin', 'ROLE::ADMIN', 'test@test.test', '$argon2id$v=19$m=16384,t=1,p=4$mpeEEsB48iz+8kanlruJXQ$nD1Pn9/H/wP8xXjgB6xpLs2Ku5jjQORkKYM2/tzruWI');
INSERT INTO public.account
( deleted_at, username, "role", email, passwd_hash)
VALUES(NULL, 'test', 'ROLE::USER', 'test@test.test', '$argon2id$v=19$m=16384,t=1,p=4$PGAqIIYQfOYquFU+fTfzXA$ArjLOdzFWgNalinA5D/BKuJwn1RyFfIRVbVUPbx1qfU');

-- for admin id should be 1
INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.1', 'ROLE::ADMIN', '*', '', '', '');

INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.1', 'ROLE::USER/OWNER', 'user.1', '', '', '');

-- for test id should be 1
INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.2', 'ROLE::USER', '*', '', '', '');

INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.2', 'ROLE::USER/OWNER', 'user.2', '', '', '');
