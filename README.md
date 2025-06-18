# blog


# Server

## 分層
1. db
1. web
2. usecase
3. service
4. domain


## 進度
### Account
- [x] Create
- [x] Login
- [x] Get
- [x] List

### Post
- [x] Create
- [x] Delete
- [x] Get
- [x] List

### Site
部落格的最小單位，與用戶脫鉤
- [x] Create
- [ ] Delete
- [ ] Get
- [x] List

### Attachtment
能與Post/Site關聯以site為擁有者單位進行附件管理
能與外部或由內部端口提供
可能需要多重關聯表
- [ ] Create
- [ ] Delete
- [ ] Get
- [ ] List

### Comment
與Post關聯
只能由匿名創建並由擁有者/管理員刪除
- [ ] Create
- [ ] Delete
- [ ] Get
- [ ] List

### AccessLog
每一項授權操作都會在系統內留下操作日誌

### db migration
https://github.com/pressly/goose