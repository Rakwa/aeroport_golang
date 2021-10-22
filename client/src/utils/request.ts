import params2query from './paramsToQuery'

export interface FetchRequestOptions {
  prefix: string
  headers: Record<string, string>
  params: Record<string, string | number | boolean>
}

export default class FetchRequest {
  private readonly defaultOptions: FetchRequestOptions = {
    prefix: '',
    headers: {},
    params: {},
  }

  private readonly options: FetchRequestOptions

  constructor(options: Partial<FetchRequestOptions> = {}) {
    this.options = Object.assign({}, this.defaultOptions, options)
  }

  private readonly generateFinalUrl = (
    url: string,
    options: Partial<FetchRequestOptions> = {}
  ): string => {
    const prefix = options.prefix ?? this.options.prefix
    const params = Object.assign({}, this.options.params, options.params ?? {})

    let finalUrl = `${prefix}${url}`
    if (Object.keys(params).length > 0) finalUrl += `?${params2query(params)}`

    return finalUrl
  }
}
