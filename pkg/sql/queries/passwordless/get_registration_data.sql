select *, credential::text as credential
FROM users
         LEFT JOIN passwordless p on users.id = p.user_id
WHERE users.id = ?
  and p.active = false
  and p.status = 'REGISTRATION_IN_PROGRESS'
ORDER BY p.created_at desc
LIMIT 1;