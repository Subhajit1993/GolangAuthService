/**
 * @name All functions
 * @description Lists all functions in the codebase for debugging purposes.
 * @kind problem
 * @problem.severity warning
 * @precision high
 * @id go/all-functions-debug
 */

import go

from Function f
select f, "Function found: " + f.getName()