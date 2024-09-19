# Role
- user
- admin


# Use case

- user
    * [ ] 1. User can create an account
    * [ ] 2. User can login
    * [ ] 3. User can logout
    * [ ] 4. User can reset password
    * [ ] 5. User can view movie database
    * [ ] 6. User can search for movies
    * [ ] 7. User can view movie details
    * [ ] 8. User can view movie reviews

- admin
    * [ ] 1. Admin can add movies
    * [ ] 2. Admin can edit movies
    * [ ] 3. Admin can delete movies
    * [ ] 4. Admin can view user details
    * [ ] 5. Admin can view user reviews
    * [ ] 6. Admin can delete user reviews
    * [ ] 7. Admin can delete user accounts

# Tech stack

- frontend: reactjs(good visual, fast, responsive)
- backend: redis (go : fast, concurrent, low latency)
- mainly store user data, movie can fetched from imdb

# Phase 1 objectives (backend: go)

- [ ] fetch movie data from imdb, return dictionary(json)
- [ ] create and store movies into database
- [ ] create facade functions

# Phase 2 objectives (frontend: reactjs)

- [ ] create movie_x website navigation
- [ ] create movie_x website hero
- [ ] create movie_x website movie list
