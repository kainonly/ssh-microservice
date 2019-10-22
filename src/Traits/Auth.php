<?php
declare(strict_types=1);

namespace Hyperf\Support\Traits;

use Hyperf\HttpServer\Contract\ResponseInterface;

/**
 * Trait Auth
 * @package Hyperf\Support\Traits
 * @property ResponseInterface $response
 */
trait Auth
{
    protected function __auth()
    {
        return $this->response->json([
            'error' => 1,
            'msg' => ''
        ]);
    }
}