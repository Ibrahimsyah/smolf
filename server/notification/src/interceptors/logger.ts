import { CallHandler, ExecutionContext, NestInterceptor } from '@nestjs/common';
import { Request, Response } from 'express';
import { Observable, tap } from 'rxjs';

export class LoggerInterceptor implements NestInterceptor {
  intercept(
    context: ExecutionContext,
    next: CallHandler<any>,
  ): Observable<any> | Promise<Observable<any>> {
    const ctx = context.switchToHttp();
    const request: Request = ctx.getRequest();
    const response: Response = ctx.getResponse();

    const timeStart = Date.now();
    return next.handle().pipe(
      tap(() => {
        console.log(
          `[${response.statusCode}] ${request.method.toUpperCase()} ${request.path} - ${Date.now() - timeStart}ms`,
        );
      }),
    );
  }
}
