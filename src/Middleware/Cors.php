<?php
declare(strict_types=1);

namespace Hyperf\Support\Middleware;

use Hyperf\Contract\ConfigInterface;
use Hyperf\Utils\Context;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\MiddlewareInterface;
use Psr\Http\Server\RequestHandlerInterface;
use Hyperf\HttpServer\Contract\ResponseInterface as HttpResponse;

/**
 * Class Cors
 * @package Hyperf\Support\Middleware
 */
class Cors implements MiddlewareInterface
{
    /**
     * @var HttpResponse
     */
    private $response;
    /**
     * @var array
     */
    private $options;

    /**
     * Cors constructor.
     * @param ContainerInterface $container
     * @param HttpResponse $response
     */
    public function __construct(ContainerInterface $container, HttpResponse $response)
    {
        $this->options = $container->get(ConfigInterface::class)->get('cors');
        $this->response = $response;
    }

    /**
     * Process an incoming server request.
     *
     * Processes an incoming server request in order to produce a response.
     * If unable to produce the response itself, it may delegate to the provided
     * request handler to do so.
     */
    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        if ($request->getMethod() == 'OPTIONS') {
            return $this->response->json([]);
        }

        $response = Context::get(ResponseInterface::class);
        if (!empty($this->options['allow_origin']) && is_array($this->options['allow_origin'])) {
            if (in_array('*', $this->options['allow_origin'])) {
                $response = $response->withHeader(
                    'Access-Control-Allow-Origin',
                    '*'
                );
            } elseif (in_array($this->options->getHeader('origin'), $this->options['allow_origin'])) {
                $response = $response->withHeader(
                    'Access-Control-Allow-Origin',
                    $request->getHeader('origin')
                );
            }
        }

        if (!empty($this->options['allow_credentials'])
            && is_bool($this->options['allow_credentials'])
            && $this->options['allow_credentials'] === true) {
            $response = $response->withHeader(
                'Access-Control-Allow-Credentials',
                'true'
            );
        }

        if (!empty($this->options['expose_headers'])
            && is_array($this->options['expose_headers'])) {
            $response = $response->withHeader(
                'Access-Control-Expose-Headers',
                implode(',', $this->options['expose_headers'])
            );
        }

        if (!empty($this->options['allow_headers'])
            && is_array($this->options['allow_headers'])) {
            $response = $response->withHeader(
                'Access-Control-Allow-Headers',
                implode(',', $this->options['allow_headers'])
            );
        }

        if (!empty($this->options['allow_headers'])
            && is_array($this->options['allow_headers'])) {
            $response = $response->withHeader(
                'Access-Control-Allow-Headers',
                implode(',', $this->options['allow_headers'])
            );
        }

        if (!empty($this->options['max_age'])
            && is_integer($this->options['max_age'])) {
            $response = $response->withHeader(
                'Access-Control-Max-Age',
                implode(',', $this->options['max_age'])
            );
        }

        Context::set(ResponseInterface::class, $response);
        return $handler->handle($request);
    }
}