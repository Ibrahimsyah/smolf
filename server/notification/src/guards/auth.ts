// Guard is basically a improved middleware to handle authentication
// and authorization since middleware does not know which function
// should be called next.

import {
  CanActivate,
  ExecutionContext,
  ForbiddenException,
  Injectable,
} from '@nestjs/common';
import { Request } from 'express';
import { Observable } from 'rxjs';

@Injectable()
export class AuthGuard implements CanActivate {
  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const request: Request = context.switchToHttp().getRequest();
    const { headers } = request;
    if (!headers['x-client-id'] || !headers['x-client-secret']) {
      throw new ForbiddenException('missing client identifier');
    }

    return true;
  }
}
