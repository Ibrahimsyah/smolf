import {
  Controller,
  ForbiddenException,
  Get,
  UseFilters,
  UseGuards,
  UseInterceptors,
} from '@nestjs/common';
import { HttpExceptionFilter } from 'src/filters/http-exception';
import { AuthGuard } from 'src/guards/auth';
import { HttpResponseInterceptor } from 'src/interceptors/http-respose';
import { HelloResponse } from './hello.type';
import { ApiHeaders, ApiTags } from '@nestjs/swagger';

@Controller()
@UseInterceptors(HttpResponseInterceptor)
@ApiTags('notification')
@ApiHeaders([
  {
    name: 'x-client-id',
    example: '76ax5ahejzt6512841a==',
  },
  {
    name: 'x-client-secret',
    example: 'asd!32aszo8917zx!',
  },
])
export class HelloController {
  @Get('/hello')
  @UseGuards(AuthGuard)
  @UseFilters(HttpExceptionFilter)
  getHello(): HelloResponse {
    return {
      id: 1,
      name: 'John doe',
    };
  }

  @Get('/hi')
  @UseFilters(HttpExceptionFilter)
  hi(): string {
    throw new ForbiddenException();
  }
}
