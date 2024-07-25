// Interceptor is basically an additional layer placed between client and backend.
// It could be client -> interceptor -> backend as request
// or backend -> interceptor -> client as response

import {
  CallHandler,
  ExecutionContext,
  Injectable,
  NestInterceptor,
} from '@nestjs/common';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable()
export class HttpResponseInterceptor implements NestInterceptor {
  intercept(
    context: ExecutionContext,
    next: CallHandler<any>,
  ): Observable<any> | Promise<Observable<any>> {
    {
      const startTime = Date.now();
      return next.handle().pipe(
        map((data) => ({
          message: 'success',
          time_elapsed: `${Date.now() - startTime}ms`,
          data,
        })),
      );
    }
  }
}
