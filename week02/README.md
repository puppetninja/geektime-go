# Q&A
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
    ```text
    ErrNoRows表示result set中row为空，应该wrap这个error抛给上层，然后由上层决定是否向用户报错或是展示结果为空。由上层做有且仅有一次的错误处理。
    ```
