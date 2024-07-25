// http exception filter is responsible to handle custom error JSON whenever
// an exception is occured. Here we can build a custom format
// for our error response

import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
} from '@nestjs/common';
import { Request, Response } from 'express';

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  catch(exception: any, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const status = exception.getStatus();

    response.status(status).json({
      message: 'error',
      statusCode: status,
      timestamp: new Date().toISOString(),
      path: request.url,
    });
  }
}
