/**
 * @name Function has too many parameters
 * @description Functions with more than 3 parameters may be confusing and difficult to use.
 * @kind problem
 * @problem.severity warning
 * @precision high
 * @id go/function-too-many-params
 * @tags maintainability
 *       readability
 */

import go

// Find all functions and methods with more than 3 parameters
from Function f
where f.getNumParameter() > 3
select f, "Function " + f.getName() + " has " + f.getNumParameter().toString() + " parameters, which exceeds the maximum of 3."