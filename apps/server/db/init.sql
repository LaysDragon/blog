INSERT INTO public.account
(id, created_at, updated_at, deleted_at, username, "role", email, passwd_hash)
VALUES(1, '2025-06-13 18:55:01.090', '2025-06-13 18:55:01.090', NULL, 'admin', 'ROLE::ADMIN', 'test@test.test', '$argon2id$v=19$m=16384,t=1,p=4$mpeEEsB48iz+8kanlruJXQ$nD1Pn9/H/wP8xXjgB6xpLs2Ku5jjQORkKYM2/tzruWI');
INSERT INTO public.account
(id, created_at, updated_at, deleted_at, username, "role", email, passwd_hash)
VALUES(2, '2025-06-03 18:06:05.249', '2025-06-03 18:06:05.249', NULL, 'test', 'ROLE::USER', 'test@test.test', '$argon2id$v=19$m=16384,t=1,p=4$PGAqIIYQfOYquFU+fTfzXA$ArjLOdzFWgNalinA5D/BKuJwn1RyFfIRVbVUPbx1qfU');

INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.1', 'ROLE::ADMIN', '*', '', '', '');

INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.1', 'ROLE::USER/OWNER', 'user.1', '', '', '');

INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.2', 'ROLE::USER', '*', '', '', '');

INSERT INTO public.casbin_rule
(p_type, v0, v1, v2, v3, v4, v5)
VALUES('p', 'user.2', 'ROLE::USER/OWNER', 'user.2', '', '', '');
