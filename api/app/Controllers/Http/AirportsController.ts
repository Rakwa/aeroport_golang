// import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

import {HttpContextContract} from "@ioc:Adonis/Core/HttpContext";

export default class AirportsController {

  public async measures({response}: HttpContextContract) {
    return response.send({message: 'Hello'})
  }
}
