// Guard is basically a improved middleware to handle authentication
// and authorization since middleware does not know which function
// should be called next.

import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { Observable } from 'rxjs';

@Injectable()
export class AuthGuard implements CanActivate {
  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    console.log(context.getType());
    return true;
  }
}
