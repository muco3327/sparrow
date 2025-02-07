package relayer

import (
	"context"

	"github.com/palomachain/sparrow/chain"
	"github.com/palomachain/sparrow/chain/paloma"
	"github.com/palomachain/sparrow/util/slice"
	log "github.com/sirupsen/logrus"
)

func (r *Relayer) Process(ctx context.Context) error {
	for chainID, p := range r.processors {
		for _, queueName := range p.SupportedQueues() {
			logger := log.WithFields(log.Fields{
				"processor-chain-id": chainID,
				"queue-name":         queueName,
			})
			// TODO: remove comments once signing is done on the paloma side.
			queuedMessages, err := r.palomaClient.QueryMessagesForSigning(ctx, queueName)
			loggerQueuedMessages := logger.WithFields(log.Fields{
				"messages": queuedMessages,
			})
			loggerQueuedMessages.Info("messages to sign")

			if err != nil {
				logger.Warn("failed getting messages to sign")
				return err
			}

			if len(queuedMessages) > 0 {
				signedMessages, err := p.SignMessages(ctx, queueName, queuedMessages...)
				if err != nil {
					loggerQueuedMessages.WithFields(log.Fields{
						"err": err,
					}).Error("unable to sign messages")
					return err
				}
				loggerQueuedMessages = loggerQueuedMessages.WithFields(log.Fields{
					"signed-messages": signedMessages,
				})
				loggerQueuedMessages.Info("signed messages")

				if err = r.broadcastSignaturesAndProcessAttestation(ctx, queueName, signedMessages); err != nil {
					loggerQueuedMessages.WithFields(log.Fields{
						"err": err,
					}).Info("couldn't broadcast signatures and process attestation")
					return err
				}
			}

			relayCandidateMsgs, err := r.palomaClient.QueryMessagesInQueue(ctx, queueName)
			if err != nil {
				logger.WithFields(log.Fields{
					"err": err,
				}).Error("couldn't get messages to relay")
				return err
			}

			logger.WithFields(log.Fields{
				"messages-to-relay": relayCandidateMsgs,
			}).Info("relaying messages")
			// if err = p.ProcessMessages(ctx, queueName, relayCandidateMsgs); err != nil {
			// 	fmt.Println("error processing a message", err)
			// 	return err
			// }
		}
	}

	return nil
}

func (r *Relayer) broadcastSignaturesAndProcessAttestation(ctx context.Context, queueTypeName string, sigs []chain.SignedQueuedMessage) error {
	broadcastMessageSignatures, err := slice.MapErr(
		sigs,
		func(sig chain.SignedQueuedMessage) (paloma.BroadcastMessageSignatureIn, error) {
			var zero paloma.BroadcastMessageSignatureIn
			var extraData []byte

			// check if this is something that requires attestation
			evidence, err := r.attestExecutor.Execute(ctx, queueTypeName, sig.Msg)
			if err != nil {
				return zero, err
			}

			if evidence != nil {
				// TODO: include evidence.Bytes() into the signature
				extraData, err = evidence.Bytes()
				if err != nil {
					return zero, err
				}
			}

			return paloma.BroadcastMessageSignatureIn{
				ID:              sig.ID,
				QueueTypeName:   queueTypeName,
				Signature:       sig.Signature,
				ExtraData:       extraData,
				SignedByAddress: sig.SignedByAddress,
			}, nil
		},
	)
	if err != nil {
		return err
	}

	return r.palomaClient.BroadcastMessageSignatures(ctx, broadcastMessageSignatures...)
}
